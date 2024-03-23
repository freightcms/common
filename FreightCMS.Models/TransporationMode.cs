namespace FreightCMS.Models;

/// <summary>
/// Represents the mode of transportation used to move goods from one location to another.
/// </summary>
/// <see cref="https://kb.freightcms.com/modes/"/> 
public enum TransportationMode
{
    Air,
    Ocean,
    Rail,
    LTL,
    FTL,
    VLTL,
    Multimodal,
    Intermodal,
    Parcel,
    Drayage,
    RollOnRollOff,
    LoadOnLoadOff,
    BreakBulk,
}