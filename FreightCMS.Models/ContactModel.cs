namespace FreightCMS.Models;

/// <summary>
/// Entity representing a contact that can be used to reach a company, organization, or individual.
/// </summary>
public class Contact
{
    /// <summary>
    /// The unique identifier for the contact.
    /// </summary>
    public required string Id { get; set; }

    /// <summary>
    /// The name of the contact.
    /// </summary>
    public required string Name { get; set; }

    /// <summary>
    /// The email address of the contact.
    /// </summary>
    public required string Email { get; set; }

    /// <summary>
    /// The phone number of the contact.
    /// </summary>
    public required string Phone { get; set; }

    /// <summary>
    /// The fax number of the contact.
    /// </summary>
    public string? Fax { get; set; }
}