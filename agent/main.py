import click
import os
import subprocess
from rich import print
from agent import generate_command
from utils import is_dangerous_command, tokenize_command

@click.command() # denotes the function as a CLI entry point
@click.argument("query", nargs=-1) # Accepts a positional argument query, allows to pass multiple words
def main(query):
    user_query = " ".join(query)
    # This check is just temporary as there is no funtionality for running the agent without a query
    if not user_query:
        print("[orange1]Please provide a query.[/orange1]")
        return
    print(f"[bold cyan]processing:[/bold cyan] {user_query}.")
    command, explanation = generate_command(user_query)
    commands = tokenize_command(command)
    commands.insert(0, "set -e")  # Ensure the script exits on error
    full_command = "\n".join(commands)
    print("[green]Query processed successfully![/green]") # This is a placeholder for the actual processing logic
    is_dangerous, warnings = is_dangerous_command(commands)
    if is_dangerous:
        print("[red]Warning: This command may be dangerous![/red]")
        for warning in warnings:
            print(f"[red]{warning}[/red]")
        print("[yellow]Do you want to continue? (y/n)[/yellow]")
        if input().lower() != "y":
            print("[red]Command execution aborted.[/red]")
            return
    print(f"[cyan]Command:[/cyan] {full_command}")
    print(f"[cyan]Explanation:[/cyan] {explanation}")
    print("[yellow]Do you want to run this command? (y/n)[/yellow]")
    if input().lower() == "y":
        try:
            process = subprocess.Popen(
                full_command,
                shell=True,
                executable="/bin/bash",
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,)
            stdout , stderr = process.communicate()
            if stdout:
                print(f"[green]Output:[/green] {stdout.decode('utf-8')}")
            if stderr:
                print(f"[red]Error:[/red] {stderr.decode('utf-8')}")
        except subprocess.CalledProcessError as e:
            print(f"[red]Error executing command: {e}[/red]")
        except Exception as e:
            print(f"[red]An unexpected error occurred: {e}[/red]")

if __name__ == "__main__":
    main()