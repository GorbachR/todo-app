package notesappapi.exception;

import java.time.LocalDateTime;
import java.util.List;
import org.springframework.http.HttpStatus;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class ErrorResponse {
    private LocalDateTime timestamp;

    private HttpStatus status;

    private String message;

    private List<String> errors;

    public ErrorResponse() {
        timestamp = LocalDateTime.now();
    }

    public ErrorResponse(HttpStatus code, String message) {
        this();
        this.status = code;
        this.message = message;
    }

    public ErrorResponse(HttpStatus code, List<String> errors, String message) {
        this(code, message);
        this.errors = errors;
    }
}