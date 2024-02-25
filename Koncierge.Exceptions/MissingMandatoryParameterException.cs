using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Exceptions
{
    public class MissingMandatoryParameterException:Exception
    {

        public MissingMandatoryParameterException() : base() { }
        public MissingMandatoryParameterException(string message) : base(message) {

            
        
        }
        public MissingMandatoryParameterException(string message, Exception innerException) : base(message, innerException) { }

        public static MissingMandatoryParameterException WithParameter(string param)
        {
            string message = $"The mandatory Parameter '{param}' has not been provided";

            return new MissingMandatoryParameterException(message);
        }
    }
}
