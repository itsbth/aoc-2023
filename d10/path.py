import re
import sys
from xml.sax.saxutils import escape

def parse_log_line(line):
    """
    Extract x and y coordinates from a log line.
    """
    match = re.search(r'x:\s*(\d+),\s*y:\s*(\d+)', line)
    if match:
        return int(match.group(1)), int(match.group(2))
    return None

def parse_enclosed_area(line):
    """
    Extract coordinates from 'Enclosed area at' log lines.
    """
    match = re.search(r'Enclosed area at (\d+),(\d+)', line)
    if match:
        return int(match.group(1)), int(match.group(2))
    return None

def create_svg_path(coordinates, colors, enclosed_areas):
    """
    Create SVG path data strings from a list of coordinates, changing colors for each new path.
    Include 1x1 squares for enclosed areas.
    """
    if not coordinates:
        return '', []
    paths = []
    squares = []
    current_path = ''
    color_index = 0

    for point in coordinates:
        if 'new_path' in point:
            if current_path:
                paths.append((colors[color_index % len(colors)], current_path))
                color_index += 1
            current_path = ''
        else:
            x, y = point
            if current_path == '':
                current_path = f'M {x},{y} '  # Start a new path
            else:
                current_path += f'L {x},{y} '  # Continue the path

    if current_path:
        paths.append((colors[color_index % len(colors)], current_path))

    for x, y in enclosed_areas:
        squares.append(f'<rect x="{x-0.5}" y="{y-0.5}" width="1" height="1" fill="black"/>')

    return paths, squares

def generate_svg(paths, squares, width=150, height=150, stroke_width=0.9):
    """
    Generate an SVG file content with multiple paths and squares for enclosed areas.
    """
    svg_paths = ''.join([f'<path d="{escape(path)}" stroke="{color}" fill="none" stroke-width="{stroke_width}"/>' for color, path in paths])
    svg_squares = ''.join(squares)
    svg_content = f'''<svg width="{512}" height="{512}" viewBox="0 0 {width} {height}" xmlns="http://www.w3.org/2000/svg">
    {svg_paths}
    {svg_squares}
</svg>
    '''
    return svg_content

# Path colors
colors = ["red", "blue", "green", "purple", "orange"]

# Read log lines from stdin
log_lines = sys.stdin.readlines()

# Extract coordinates and identify new paths and enclosed areas
coordinates = []
enclosed_areas = []
max_x = max_y = 0

for line in log_lines:
    if "Trying dir" in line:
        coordinates.append(('new_path', 'new_path'))
    elif "Enclosed area at" in line:
        enclosed_area = parse_enclosed_area(line)
        if enclosed_area:
            enclosed_areas.append(enclosed_area)
            max_x = max(max_x, enclosed_area[0])
            max_y = max(max_y, enclosed_area[1])
    else:
        coord = parse_log_line(line)
        if coord:
            coordinates.append(coord)
            max_x = max(max_x, coord[0])
            max_y = max(max_y, coord[1])

# Create SVG paths and squares
paths, squares = create_svg_path(coordinates, colors, enclosed_areas)

# Generate SVG content
svg_content = generate_svg(paths, squares, width=max_x + 10, height=max_y + 10)

# Write to an SVG file
with open("path.svg", "w") as file:
    file.write(svg_content)

print("SVG file generated successfully.")
