using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Core.Exceptions
{
    public class KubeConfigNotFoundException : Exception
    {
        public KubeConfigNotFoundException(string kcId) : base($"Kubeconfig '{kcId}' does not exists.") { }
    }

    public class PodForServiceNotFoundException : Exception
    {
        public PodForServiceNotFoundException(string service) : base($"A pod for '{service}' cannot be found.") { }
    }
}
