# Grok Code

![](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Grok Code is a maximalist, open-source agentic coding tool that lives right in your terminal. With strong "Grok vibes", it understands your codebase, plans complex tasks, edits files, runs tests, and handles your git workflows using natural language commands. 

Unlike other tools, Grok Code features a **smart multi-model router**, seamlessly switching between xAI (Grok) for deep reasoning, DeepSeek for cost-effective operations, and Gemini for sheer speed.

**Learn more in our [official documentation](./docs/overview.md)** (Coming soon).

<img src="./demo.gif" alt="Grok Code Demo" />

## Get Started

> [!NOTE]
> Pre-compiled binaries are the fastest way to get started. No Go installation required.

1. Install Grok Code:

    **MacOS/Linux (Recommended):**
    ```bash
    curl -fsSL https://raw.githubusercontent.com/holasoymalva/grok-code/main/install.sh | bash
    ```

    **Homebrew (MacOS/Linux):**
    ```bash
    brew install holasoymalva/tap/grok-code
    ```

    **Go Install (Developers):**
    ```bash
    go install github.com/holasoymalva/grok-code/cmd/grokcode@latest
    ```

2. Run `grokcode chat` to initialize the wizard and configure your API keys.

## Core Features

- **Multi-Model Support**: Use xAI (Grok), DeepSeek, Gemini, Groq, or even local LLMs (like Ollama).
- **Smart Model Router**: Dynamically route tasks. Let DeepSeek handle simple scaffolding while Grok tackles complex architectural decisions.
- **Agentic Loop**: Built-in capabilities for planning, tool execution, interactive diffs, and self-correction.
- **Rich TUI**: Powered by Bubble Tea and Lipgloss for a dynamic, visually stunning terminal experience.

## Reporting Bugs

We welcome your feedback and code contributions. Since Grok Code is fully open source, you can inspect everything under the hood. Open a [GitHub issue](https://github.com/holasoymalva/grok-code/issues) or submit a Pull Request.

## Connect with the Community

Join the [Grok Code Developers Discord](#) (Link coming soon) to connect with other developers. Get help, share your own custom tools, and discuss the future of open-source agentic coding.

## Data Privacy & Telemetry

Grok Code is fully open source and runs entirely on your local machine. 

**We do not collect telemetry, code snippets, or usage data.** Your code stays on your machine, and your prompts are only sent directly to the model providers you explicitly configure. There are no middlemen.
