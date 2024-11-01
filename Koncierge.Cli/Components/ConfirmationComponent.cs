using Spectre.Console;

namespace Koncierge.Cli.Components;

public static class ConfirmationComponent
{
    


    public static bool Ask(string question, bool defaultAnswer=true)
    {
        
       return AnsiConsole.Prompt(
            new TextPrompt<bool>(question)
                .AddChoice(true)
                .AddChoice(false)
                .DefaultValue(defaultAnswer)
                .WithConverter(choice => choice ? "y" : "n"));
        
        
        
    }
}