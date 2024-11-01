using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Domain.Enums
{
    public enum KonciergeConnectionStatus
    {
        Connected,
        Disconnected,
        Failed,
        NotFound


    }

    public enum KonciergeForwardType
    {
        Service,
        Pod


    }
    
    public enum KonciergeConfigType
    {
        ConfigMap,
        Secret


    }

  
}
