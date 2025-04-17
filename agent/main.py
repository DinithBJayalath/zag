import click
from rich import print
from agent import generate_command

@click.command() # denotes the function as a CLI entry point
@click.argument("query", nargs=-1) # Accepts a positional argument query, allows to pass multiple words
def main(query):
    user_query = " ".join(query)
    # This check is just temporary as there is no funtionality for running the agent without a query
    if not user_query:
        print("[orange1]Please provide a query.[/orange1]")
        return
    print(f"[cyan]processing:[/cyan] {user_query}.")
    command, explanation = generate_command(user_query)
    print("[green]Query processed successfully![/green]") # This is a placeholder for the actual processing logic
    print(f"[cyan]Command:[/cyan] {command}")
    print(f"[cyan]Explanation:[/cyan] {explanation}")

if __name__ == "__main__":
    main()