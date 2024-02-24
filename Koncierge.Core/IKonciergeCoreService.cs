using Koncierge.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Core
{
    public interface IKonciergeCoreService
    {
        public Task<List<KubeConfigFile>> GetKubeConfigFromPath(bool dry=false);
        public Task<List<KubeConfigFile>> GetKubeConfigFromPath(string path, bool dry = false);

        public string GetKubeConfigDefaultPath();
    }
}
