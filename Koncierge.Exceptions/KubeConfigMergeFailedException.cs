using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Exceptions
{
    public class KubeConfigMergeFailedException:Exception
    {

        public KubeConfigMergeFailedException() : base() { }
        public KubeConfigMergeFailedException(string message) : base(message) { }
        public KubeConfigMergeFailedException(string message, Exception innerException) : base(message, innerException) { }

        public static KubeConfigMergeFailedException WithFromTo(string from, string to)
        {
            string message = $"Error Merging {from} into {to}";

            return new KubeConfigMergeFailedException(message);
        }
    }
}
