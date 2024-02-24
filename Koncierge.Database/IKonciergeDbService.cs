using Koncierge.Database.Repositories;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Database
{
    public interface IKonciergeDbService:IDisposable
    {
        public IKubeConfigFileRepository KubeConfigFileRepository();
    }
}
