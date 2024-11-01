using System.Linq.Expressions;
using Koncierge.Domain.Entities;
using Spectre.Console;

namespace Koncierge.Cli.Components;

public static class MultiselectComponent<T> where T : notnull
{


    public static T One(IQueryable<T> items, string entityName, int size = 10)
    {
        T ret = default!;
        
        if (items.Count() > 1)
        {
            ret = AnsiConsole.Prompt(
                new SelectionPrompt<T>()
                    .Title($"{Emoji.Known.MagnifyingGlassTiltedLeft} Select the [green]{entityName}[/] ")
                    .PageSize(size)
                    .MoreChoicesText($"[grey](Move up and down to reveal more {entityName}(s))[/]")
                    .AddChoices(items)
                    

                    .EnableSearch()
            );
        }
        else if (items.Count() == 1)
        {
            ret = items.First();

        }

        return ret;
    }


    public static List<T> Many(IQueryable<T> items, string entityName, int size = 10, bool required=true)
    {
        List<T> ret = default!;
        
        if (items.Count() > 1)
        {
            ret = AnsiConsole.Prompt(
                new MultiSelectionPrompt<T>()
                    .Title($"{Emoji.Known.MagnifyingGlassTiltedLeft} Select one or more [green]{entityName}[/] ")
                    .PageSize(size)
                    .MoreChoicesText($"[grey](Move up and down to reveal more {entityName}(s))[/]")
                    .AddChoices(items).Required(required)

                    
            );
        }
        else if (items.Count() == 1)
        {
            ret = items.ToList();

        }

        return ret;
    }


}