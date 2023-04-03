#!/usr/bin/env python3
import subprocess

def wrap(text, limit, force='', start='  '):
    tokens = text.split()
    lines = [start]
    while len(tokens) > 0:
        if len(lines[-1]) + len(' ' + tokens[0]) > limit or any((tokens[0].endswith(c) for c in force)):
            lines.append(start)
        if lines[-1] != start:
            lines[-1] += ' '
        lines[-1] += tokens[0]
        tokens = tokens[1:]
    return '\n'.join(lines)

if __name__ == "__main__":
    subprocess.run(["go", "build"], check=True)
    p = subprocess.run(["./protoc-gen-protobufjs", "-help"], capture_output=True, encoding='utf-8')
    lines = p.stderr.split('\n')
    flags = []
    for line in lines:
        if line.startswith('    '):
            desc = line.strip()
            flags.append((name, desc))
        elif line.startswith('  '):
            name = line.strip()

    flags.sort(key=lambda x: 1 if 'rare' in x[1] else 0)

    lines = []
    for (name, desc) in flags:
        name_parts = name.split(' ')
        lines.append(name_parts[0] + ': ' + " ".join(name_parts[1:]))
        lines.append(wrap(desc, 74, start='  ', force=''))
        lines.append('')

    with open('README.md', 'r') as f:
        readme_lines = f.readlines()
    section_start = readme_lines.index('<!-- #help:start -->\n')
    section_end = readme_lines.index('<!-- #help:end -->\n')

    readme_lines[section_start+3:section_end-2] = [line + '\n' for line in lines]
    print(''.join(readme_lines))

    with open('README.md', 'w') as f:
        f.writelines(readme_lines)
