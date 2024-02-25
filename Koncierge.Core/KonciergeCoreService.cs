using Koncierge.Database;
using Koncierge.KubeConfig;
using Koncierge.Models;
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
        private readonly IKonciergeDbService _konciergeDbService;

        public KonciergeCoreService(IKonciergeDbService kdbs) {

            _kubeConfig=new KubeConfigService();
            _konciergeDbService = kdbs;


        }

       
    }
}
