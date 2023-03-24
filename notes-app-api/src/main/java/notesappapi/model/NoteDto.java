package notesappapi.model;

import java.sql.Timestamp;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Size;
import lombok.Value;

@Value
public class NoteDto {
    long id;

    @NotNull(message = "title can't be null")
    @Size(max = 50, message = "title can't have more than 50 chars")
    String title;

    @NotNull(message = "body can't be null")
    @Size(max = 500, message = "body can't have more than 500 chars")
    String body;
    Timestamp lastEditedTimestamp;
}