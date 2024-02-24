using Koncierge.Models.Config;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.KubeConfig
{
    public interface  IKubeConfigService
    {

        public string GetKubeConfigDefaultPath();
        public Task<List<KubeConfigFile>> GetKubeConfigFromPath();
        public Task<List<KubeConfigFile>> GetKubeConfigFromPath(string p);


        
    }
}
