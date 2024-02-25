using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using static Koncierge.Models.Enums;

namespace Koncierge.Models
{
    public class KubeConfigFile: BaseEntity
    {
        public string Name { get; set; }
        public string Path { get; set; }
        public DateTime Added { get; set; }
        public DateTime? Modified { get; set; }

        public KubeConfigFileStatus Status { get; set; }
        public bool JustAdded { get; set; }

        public KubeConfigFile()
        {

            
        }
        public KubeConfigFile(string name, string path)
        {

            Name = name;
            Path = path;
            Status = KubeConfigFileStatus.OK;
            Added = DateTime.Now;
            JustAdded = true;
        }




    }
}
