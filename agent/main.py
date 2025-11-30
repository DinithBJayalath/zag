from agent import generate_command
from utils import is_dangerous_command, tokenize_command

def processCommand(query):
    user_query = " ".join(query)
    if not user_query:
        return "No natural language query found"
    command, explanation = generate_command(user_query)
    commands = tokenize_command(command)
    commands.insert(0, "set -e")  # Ensure the script exits on error
    full_command = "\n".join(commands)
    is_dangerous, warnings = is_dangerous_command(commands)
    return full_command, explanation, is_dangerous