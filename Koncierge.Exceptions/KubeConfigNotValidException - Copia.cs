using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Exceptions
{
    public class KubeConfigNotSavedException:Exception
    {

        public KubeConfigNotSavedException() : base() { }
        public KubeConfigNotSavedException(string message) : base(message) { }
        public KubeConfigNotSavedException(string message, Exception innerException) : base(message, innerException) { }

        public static KubeConfigNotSavedException WithKcf(string kcf)
        {
            string message = $"The KubeConfig File '{kcf}' has not been saved.";

            return new KubeConfigNotSavedException(message);
        }
    }
}
