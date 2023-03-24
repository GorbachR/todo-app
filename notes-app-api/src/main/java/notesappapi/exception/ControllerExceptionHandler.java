package notesappapi.exception;

import java.util.List;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

@RestControllerAdvice
public class ControllerExceptionHandler extends ResponseEntityExceptionHandler {

    @Override
    protected ResponseEntity<Object> handleMethodArgumentNotValid(MethodArgumentNotValidException e,
            HttpHeaders headers, HttpStatusCode status, WebRequest request) {

        HttpStatus returnStatus = HttpStatus.valueOf(status.value());
        List<String> errors = e.getBindingResult().getFieldErrors().stream()
                .map(err -> err.getRejectedValue() + " invalid, " + err.getDefaultMessage())
                .distinct()
                .toList();

        return new ResponseEntity<Object>(new ErrorResponse(returnStatus, errors, "Arguments invalid"), returnStatus);
    }
}