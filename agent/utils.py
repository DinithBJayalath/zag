def normalize_command(cmd):
    """
    Takes a multi-line command string and splits it into individual shell commands.
    Ignores blank lines, removes comment lines and trims each command.
    """
    lines = cmd.strip().splitlines()
    commands = [line.strip() for line in lines if line.strip() and not line.strip().startswith("#")]
    return commands