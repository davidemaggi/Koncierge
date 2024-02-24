using Koncierge.KubeConfig;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Core
{
    public partial class KonciergeCoreService: IKonciergeCoreService
    {
        private readonly IKubeConfigService _kubeConfig;

        public KonciergeCoreService() {

            _kubeConfig=new KubeConfigService();



        }


    }
}
