using Koncierge.Core.Services.Implementations;
using Koncierge.Core.Services.Interfaces;
using Koncierge.Domain;
using Koncierge.Domain.Repositories.Implementations;
using Koncierge.Domain.Repositories.Interfaces;
using Microsoft.Extensions.Logging;

namespace Koncierge.Gui
{
    public static class MauiProgram
    {
        public static MauiApp CreateMauiApp()
        {
            var builder = MauiApp.CreateBuilder();
            builder
                .UseMauiApp<App>()
                .ConfigureFonts(fonts =>
                {
                    fonts.AddFont("OpenSans-Regular.ttf", "OpenSansRegular");
                });

            builder.Services.AddMauiBlazorWebView();

#if DEBUG
    		builder.Services.AddBlazorWebViewDeveloperTools();
    		builder.Logging.AddDebug();
#endif

            builder.Services.AddLogging();
            builder.Services.AddDbContext<KonciergeContext>();


            builder.Services.AddSingleton<IKubeConfigRepository, KubeConfigRepository>();
            builder.Services.AddSingleton<IKonciergeService, KonciergeService>();


#if DEBUG



            AppDomain.CurrentDomain.FirstChanceException += (s, e) =>
            {
                System.Diagnostics.Debug.WriteLine("********** FE Exception **********");
                System.Diagnostics.Debug.WriteLine(e.Exception);
            };
#endif

            //var dbContext = new KonciergeContext();
            //dbContext.Initialize();
           

            return builder.Build();
        }
    }
}
