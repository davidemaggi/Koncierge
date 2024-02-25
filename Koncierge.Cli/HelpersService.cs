using Spectre.Console;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Cli
{
    public class HelpersService
    {





        public void WriteInfo(string msg) => AnsiConsole.MarkupLine($":information: {msg}");
        public void WriteError(string msg) => AnsiConsole.MarkupLine($":no_entry: {msg}");
        public void WriteWarning(string msg) => AnsiConsole.MarkupLine($":warning: {msg}");
        public void WriteSuccess(string msg) => AnsiConsole.MarkupLine($":check_mark_button: {msg}");
        public void WriteFatal(string msg) => AnsiConsole.MarkupLine($":skull: {msg}");
        public void WriteOk(string msg) => AnsiConsole.MarkupLine($":thumbs_up: {msg}");
        public void WriteKoncierge(string msg) => AnsiConsole.MarkupLine($":person_in_tuxedo: {msg}");


        







    }
}
