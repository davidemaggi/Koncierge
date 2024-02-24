using Koncierge.Cli.Commands;
using Koncierge.Cli.Commands.KubeConfig;
using Koncierge.Cli.Infrastructure;
using Koncierge.Core;

using Microsoft.Extensions.DependencyInjection;
using Spectre.Console;
using Spectre.Console.Cli;

namespace Koncierge.Cli
{
    internal class Program
    {
        static int Main(string[] args)
        {

            Console.OutputEncoding = System.Text.Encoding.UTF8;
            var services = new ServiceCollection();

            ConfigureServices(services);

            var registrar = new TypeRegistrar(services);
            // Create a new command app with the registrar


            var app = new CommandApp<InfoCommand>(registrar);

            app.Configure(config =>
            {
                config.SetApplicationName("koncierge");

                config.AddBranch<KubeConfigSettings>("config", add =>
                {
                    add.AddCommand<GetKubeConfigCommand>("get");
                    //add.AddCommand<AddReferenceCommand>("reference");
                });

                //    config.AddCommand<WizardCommand>("wizard")
                //    .WithAlias("w")
                //    .WithDescription("Set your new Context and Namespace");

                //    config.AddCommand<MergeCommand>("merge")
                //   .WithAlias("m")
                //   .WithDescription("Merge new YAML file with your KubeConfig");

                //    config.AddCommand<ForwardCommand>("forward")
                //   .WithAlias("f")
                //   .WithAlias("fwd")
                //   .WithDescription("Forward a service/pod port to localhost");
                //    config.AddCommand<AliasCommand>("alias")
                //.WithAlias("a")
                //.WithDescription("Change Context Name to an alias");

                //    config.AddCommand<DeleteCommand>("delete")
                //   .WithAlias("d")
                //   .WithDescription("Delete context from KubeConfig");

                config.AddCommand<InfoCommand>("info")
              .WithAlias("i")
              .WithAlias("v")
              .WithAlias("version")
              .WithDescription("Get info about the current configuration");


                config.SetExceptionHandler((ex) =>
                {
                    AnsiConsole.WriteException(ex, ExceptionFormats.ShortenEverything);
                });

#if DEBUG

                config.ValidateExamples();

#endif
            });



            try
            {
                app.Run(args);
                return 0;
            }
            catch (Exception ex)
            {
                AnsiConsole.WriteException(ex, ExceptionFormats.ShortenEverything);
                return -99;
            }
        }


        public static void ConfigureServices(IServiceCollection services)
        {


            services.AddSingleton<IKonciergeCoreService, KonciergeCoreService>();
            var registrar = new TypeRegistrar(services);



        }


    }
}
