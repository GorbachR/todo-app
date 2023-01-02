package notesappapi.exception;

public class InvalidCursorException extends RuntimeException {
    
    public InvalidCursorException() {
        super("Cursor is invalid");
    }
}
