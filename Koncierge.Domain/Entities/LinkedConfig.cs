using Koncierge.Domain.Enums;

namespace Koncierge.Domain.Entities;

public class LinkedConfig: BaseEntity
{
    public KonciergeConfigType Type { get; set; }
    public string Name { get; set; }
    public List<string>  Values { get; set; }
}