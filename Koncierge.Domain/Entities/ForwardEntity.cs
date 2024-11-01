using Koncierge.Domain.Enums;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Domain.Entities
{
    public class ForwardEntity : BaseEntity
    {

        public string Context { get; set; }


        public KonciergeForwardType Type { get; set; }

        public string Selector { get; set; }

        public int LocalPort { get; set; }
        public int RemotePort { get; set; }


        public KubeConfigEntity WithConfig { get; set; }
        public string Namespace { get; set; }
        
        public ICollection<LinkedConfig> Configs { get; set; }
        
        
        public override string ToString()
        {

            if (Id==Guid.Empty) { 
            
                return this.Context;

            }


            return $"{this.Selector}:{this.LocalPort} --> localhost:{this.RemotePort}";
        }
        
        
        
    }
}
