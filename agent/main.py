import click
from rich import print

@click.command() # denotes the function as a CLI entry point
@click.argument("query", nargs=-1) # Accepts a positional argument query, allows to pass multiple words
def main(query):
    user_query = " ".join(query)
    # This check is just temporary as there is no funtionality for running the agent without a query
    if not user_query:
        print("[orange1]Please provide a query.[/orange1]")
        return
    print(f"[cyan]processing:[/cyan] {user_query}.")
    print("[green]Query processed successfully![/green]") # This is a placeholder for the actual processing logic

if __name__ == "__main__":
    main()