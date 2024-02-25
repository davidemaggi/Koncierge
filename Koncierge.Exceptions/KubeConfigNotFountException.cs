using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Exceptions
{
    public class KubeConfigNotFountException:Exception
    {

        public KubeConfigNotFountException() : base() { }
        public KubeConfigNotFountException(string message) : base(message) { }
        public KubeConfigNotFountException(string message, Exception innerException) : base(message, innerException) { }

        public static KubeConfigNotFountException WithKcf(string kcf)
        {
            string message = $"The KubeConfig File '{kcf}' has not been found";

            return new KubeConfigNotFountException(message);
        }
    }
}
