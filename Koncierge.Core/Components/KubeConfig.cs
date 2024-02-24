using Koncierge.Models.Config;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Core
{
    public partial class KonciergeCoreService
    {

        public async Task<List<KubeConfigFile>> GetKubeConfigFromPath(bool dry = false) { 
            
            
            var configs= await _kubeConfig.GetKubeConfigFromPath();

            if (!dry) { SaveConfigs(configs); }

            return configs;
        }


        public async Task<List<KubeConfigFile>> GetKubeConfigFromPath(string path, bool dry = false)  {


            var configs = await _kubeConfig.GetKubeConfigFromPath(path); 
        
            if (!dry) { SaveConfigs(configs); }

            return configs;
        }

        public string GetKubeConfigDefaultPath() => _kubeConfig.GetKubeConfigDefaultPath();

        private List<KubeConfigFile> SaveConfigs(List<KubeConfigFile> configs) {


            return configs;
        }




    }
}
