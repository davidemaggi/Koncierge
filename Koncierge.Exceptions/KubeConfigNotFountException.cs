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


    }
}
