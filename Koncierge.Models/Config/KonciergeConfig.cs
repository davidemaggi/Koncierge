using System;
using System.Collections.Generic;
using System.Linq;
using System.Security.AccessControl;
using System.Text;
using System.Threading.Tasks;
using static Koncierge.Models.Enums;

namespace Koncierge.Models.Config
{
    public class KonciergeConfig
    {

        public List<KubeConfigFile> KnownKubeConfigs { get; set; } = new List<KubeConfigFile>();


    }

    public class KubeConfigFile
    {
        public string Name { get; set; }
        public string Path { get; set; }
        public DateTime Added { get; set; }
        public DateTime? Updated { get; set; }

        public KubeConfigFileStatus Status { get; set; }


        public KubeConfigFile(string name, string  path) { 
        
            Name = name;
            Path = path;
            Status = KubeConfigFileStatus.OK;
            Added = DateTime.Now;
        
        }




    }


}
