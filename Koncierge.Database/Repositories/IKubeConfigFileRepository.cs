using Koncierge.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Numerics;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Database.Repositories
{
    public interface IKubeConfigFileRepository : IBaseRepository<KubeConfigFile>
    {
        public void ClearJustAdded();
    }
}
