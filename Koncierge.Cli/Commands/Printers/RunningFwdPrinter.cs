using Koncierge.Domain.DTOs;
using Koncierge.Domain.Entities;
using Koncierge.Domain.Enums;
using Spectre.Console;

namespace Koncierge.Cli.Commands.Printers;

public static class RunningFwdPrinter
{


    public static void Print(ForwardEntity forwardEntity, KubeConfigDto? mapsAndSecrets)
    {

        
        // Create a table
        var table = new Table();
        table.HideHeaders();

// Add some columns
        table.AddColumn(new TableColumn("label").LeftAligned());
        table.AddColumn(new TableColumn("content").LeftAligned());

// Add some rows

      //  var path = new TextPath(forwardEntity.WithConfig.Path);


      table.AddRow("Status", $"{Emoji.Known.CheckMark}[{Color.Grey}] Forwarding[/]");
      table.AddRow("Context", $"{forwardEntity.Context}[{Color.Grey}]@{forwardEntity.WithConfig.Path}[/]");
      
      var icon=forwardEntity.Type == KonciergeForwardType.Service ? Emoji.Known.GlobeWithMeridians : Emoji.Known.Package;


      
      table.AddRow("Forwarding", $"{icon} [{Color.Grey}] {forwardEntity.Namespace}.[/]{forwardEntity.Selector}:[{Color.DeepSkyBlue1}]{forwardEntity.RemotePort}[/] --> localhost:[{Color.DeepSkyBlue1}]{forwardEntity.LocalPort}[/]");
 
      var tableDetails = new Table();
      tableDetails.Title = new TableTitle("Details");
      tableDetails.HideHeaders();
      tableDetails.AddColumn(new TableColumn("label").LeftAligned());
      tableDetails.AddColumn(new TableColumn("label").LeftAligned());

      tableDetails.AddRow(buildTable("️️ConfigMaps", forwardEntity.Configs.Where(x=>x.Type==KonciergeConfigType.ConfigMap),mapsAndSecrets), buildTable("Secrets",forwardEntity.Configs.Where(x=>x.Type==KonciergeConfigType.Secret),mapsAndSecrets));

// Render the table to the console
        AnsiConsole.Write(table);
        if (forwardEntity.Configs.Any())
        {
            AnsiConsole.Write(tableDetails);

        }

        
        
    }

    private static Table buildTable(string title, IEnumerable<LinkedConfig> toPrint, KubeConfigDto? mapsAndSecrets)
    {
        var ret = new Table();
       

        ret.Title = new TableTitle(title);
        ret.HideHeaders();
        var linkedConfigs = toPrint.ToList();
        if (linkedConfigs.Any())
        {
            ret.AddColumn(new TableColumn("label").LeftAligned());
            ret.AddColumn(new TableColumn("content").LeftAligned());
            foreach (var linkedConfig in linkedConfigs)
            {
                foreach (var linkedConfigValue in linkedConfig.Values)
                {
                    ret.AddRow(linkedConfigValue, getValue(linkedConfig.Name,linkedConfigValue, linkedConfig.Type, mapsAndSecrets));

                }
            }
        }
        else
        {
            ret.AddColumn(new TableColumn("label").LeftAligned());
            ret.AddRow($"No linked {title}");

        }
        

        return ret;
    }


    private static string getValue(string main, string inner, KonciergeConfigType type, KubeConfigDto? mapsAndSecrets)
    {
        if (mapsAndSecrets is not null)
        {
            if (type==KonciergeConfigType.Secret && mapsAndSecrets.Secrets.Any())
            {
                var tmp=mapsAndSecrets.Secrets.FirstOrDefault(x=>x.Selector==main);

                if (tmp is not null)
                {
                    var tmpVal = "";
                    var hasval = tmp.Values.TryGetValue(inner,out tmpVal);
                    if (hasval)
                    {
                        return tmpVal;
                    }
                }


            }
            if (type==KonciergeConfigType.ConfigMap && mapsAndSecrets.ConfigMaps.Any())
            {
                var tmp=mapsAndSecrets.ConfigMaps.FirstOrDefault(x=>x.Selector==main);

                if (tmp is not null)
                {
                    var tmpVal = "";
                    var hasval = tmp.Values.TryGetValue(inner,out tmpVal);
                    if (hasval)
                    {
                        return tmpVal;
                    }
                }


            }
        }
        
        
        
        return "-";
    }

}