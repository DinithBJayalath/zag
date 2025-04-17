import click
import os
import subprocess
from rich import print
from agent import generate_command
from utils import normalize_command

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
    commands = normalize_command(command)
    full_command = "\n".join(commands)
    print("[green]Query processed successfully![/green]") # This is a placeholder for the actual processing logic
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