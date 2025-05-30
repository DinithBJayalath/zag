import re

DANGEROUS_COMMAND_PATTERNS = [
    r"rm\s+-rf\s+/",                    # wipes root directory
    r"rm\s+-rf\s+\*",                   # wipes all in directory
    r"mkfs\.",                          # formats disk
    r"dd\s+if=",                        # raw disk writing
    r":\(\)\s*{\s*:\s*\|\s*:\s*;\s*}",  # fork bomb
    r"shutdown\b|reboot\b",             # system shutdowns
    r"curl\s+[^\|]+\|\s*sh",            # downloading and piping to shell
    r"chmod\s+777",                     # insecure permissions
    r"yes\b",                           # infinite loop
    r"sudo\s+[^ ]*",                    # sudo without awareness
]

def split_multi_line_commands(cmd):
    """
    Takes a multi-line command string and splits it into individual shell commands.
    Ignores blank lines, removes comment lines and trims each command.
    """
    lines = cmd.strip().splitlines()
    commands = [line.strip() for line in lines if line.strip() and not line.strip().startswith("#")]
    return commands

def is_dangerous_command(command):
    """
    Checks if the command contains any dangerous patterns.
    Returns True if it does, False otherwise.
    """
    warnings = []
    for pattern in DANGEROUS_COMMAND_PATTERNS:
        if re.search(pattern, command):
            warnings.append(f"Warning: {pattern}")
    if warnings:
        return True, warnings
    return False, []