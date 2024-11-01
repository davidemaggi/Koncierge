using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using static System.Environment;

namespace Koncierge.Core.Tools
{
    public static class FileSystemTools
    {

        public static string getKonciergePath() => Path.Combine(Environment.GetFolderPath(SpecialFolder.LocalApplicationData, SpecialFolderOption.DoNotVerify), "koncierge");

    }
}
