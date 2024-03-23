namespace FreightCMS.Models;

/// <summary>
/// Represents a reference to a model that can be used to identify a company, organization, or individual.
/// Typically this is when a blind entity needs to reference their information but make it found only through
/// possible reference data known to them such as an Order Number, tracking number, purchase order, shipment number, 
/// Bill of lading, etc.
/// </summary>
public class ReferenceModel
{
    /// <summary>
    /// Unique identifier for the reference for internal usage only.
    /// </summary>
    public required string Id { get; set; }
    /// <summary>
    /// name or code to relay to the consumer of what the reference may be for.
    /// </summary>
    public required string NameOrCode { get; set; }
    /// <summary>
    /// Description of the reference that tells more about the name or code that is implied.
    /// </summary>
    public string? Description { get; set; }

    /// <summary>
    /// Any data that can be used by the end user to find information. JSON, XML, URLs, etc.
    public required string Value { get; set; }
}