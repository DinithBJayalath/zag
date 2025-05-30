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

def tokenize_command(full_command):
    """
    Tokenizes a shell command into its components.
    Returns a list of tokens.
    """
    tokens = []
    # Splitting multi-line commands into separate commands
    command_lines = full_command.strip().splitlines()
    commands = [line.strip() for line in command_lines if line.strip() and not line.strip().startswith("#")]
    for command in commands:
        temp_tokens = re.split(r'\s*(?;|&&|\|\||;|&)\s*', command) # split returns a list, so save temporarily
        for token in temp_tokens: # iterate over the temporary lisst and add to the main list
            tokens.append(token)
    return [token.strip() for token in tokens if token.strip()]