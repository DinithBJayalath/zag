import openai
from dotenv import load_dotenv
import os
import re
from prompt_template import command_prompt

load_dotenv(override=True)
openai.api_key = os.getenv("OPENAI_API_KEY")

LLM_client = openai.OpenAI()

system_message = """
You are a terminal assistant. When the user asks for help, return a shell script that could also be multi-step, inside a bash code block (```bash).
Include only the commands inside the code block. After the block, add a plain-text explanation.
"""

def generate_command(query):
    prompt = command_prompt.format(query=query)
    response = LLM_client.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "system", "content": system_message},
            {"role": "user", "content": prompt}
        ],
        temperature=0,
        max_tokens=150,
    )
    result = response.choices[0].message.content.strip()
    command, explanation = parse_response(result)
    return command.strip(), explanation.strip() if explanation else "No explanation provided."

def parse_response(content):
    """
    Extracts command and explanation from LLM response.
    Supports code blocks and plain text fallbacks.
    """
    # Step 1: Try to extract from code block
    code_block_match = re.search(r"```(?:bash)?\s*(.*?)```", content, re.DOTALL)
    if code_block_match:
        command = code_block_match.group(1).strip()
        explanation = content.replace(code_block_match.group(0), "").strip()
        return command, explanation

    # Step 2: Try to split on first newline if no code block
    lines = content.strip().split("\n", 1)
    command = lines[0].strip()
    explanation = lines[1].strip() if len(lines) > 1 else "No explanation provided."
    return command, explanation