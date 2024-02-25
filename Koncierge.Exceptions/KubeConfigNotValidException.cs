using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Exceptions
{
    public class KubeConfigNotValidException:Exception
    {

        public KubeConfigNotValidException() : base() { }
        public KubeConfigNotValidException(string message) : base(message) { }
        public KubeConfigNotValidException(string message, Exception innerException) : base(message, innerException) { }

        public static KubeConfigNotValidException WithKcf(string kcf)
        {
            string message = $"The KubeConfig File '{kcf}' is not valid.";

            return new KubeConfigNotValidException(message);
        }
    }
}
