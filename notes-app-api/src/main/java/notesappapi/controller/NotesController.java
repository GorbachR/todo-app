package notesappapi.controller;

import java.util.List;
import org.apache.commons.codec.binary.Base64;
import org.apache.commons.lang3.math.NumberUtils;
import org.springframework.data.domain.PageRequest;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.util.UriComponents;
import org.springframework.web.util.UriComponentsBuilder;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.validation.Valid;
import notesappapi.entity.Note;
import notesappapi.exception.InvalidCursorException;
import notesappapi.exception.NoteNotFoundException;
import notesappapi.repository.NotesRepository;


@RestController
@RequestMapping(path = "/api/notes")
public class NotesController {

    private final NotesRepository notesRepository;

    public NotesController(NotesRepository notesRepository) {
        this.notesRepository = notesRepository;
    }


    @GetMapping
    public ResponseEntity<List<Note>> getNotes(@RequestParam(defaultValue = "") String cursor
    , @RequestParam(defaultValue = "3") int limit, HttpServletRequest request) {
        
        HttpHeaders headers = new HttpHeaders();
        UriComponents responseLink;
        long startIndex = 0;


        if(cursor.length() > 0) {

            if(!Base64.isBase64(cursor.getBytes())) throw new InvalidCursorException();

            String decodedCursor = new String(Base64.decodeBase64(cursor));

            if(!NumberUtils.isParsable(decodedCursor)) throw new InvalidCursorException();
            
            startIndex = Long.parseLong(decodedCursor);

        }
        else startIndex = notesRepository.count();
        
        List<Note> body = notesRepository.findByIdLessThanEqualOrderByIdDesc(startIndex, PageRequest.ofSize(limit));


        if(!notesRepository.findByIdLessThanEqualOrderByIdDesc(startIndex - (long) limit, PageRequest.ofSize(limit)).isEmpty()) {

            String nextCursor = Long.toString(startIndex - (long) limit);
            nextCursor = Base64.encodeBase64URLSafeString(nextCursor.getBytes());

            // Could be a problem when reverse proxy/ load balancer is in use
            responseLink = UriComponentsBuilder
            .fromHttpUrl(request.getRequestURL().toString())
            .queryParam("cursor", nextCursor).build();

            headers.set(HttpHeaders.LINK, responseLink.toString());
        }

        return ResponseEntity.ok().headers(headers).body(body);
    }


    @PostMapping(consumes = {"application/json"})
    public ResponseEntity<Note> addNote(@Valid @RequestBody Note newNote) {
        
        notesRepository.save(newNote);
        return ResponseEntity.status(HttpStatus.CREATED).body(newNote);
    }

    @PutMapping(path = "/{id}", consumes = {"application/json"})
    public ResponseEntity<Note> editNote(@PathVariable long id, @Valid @RequestBody Note updatedNote) {

        Note note = notesRepository.findById(id)
        .orElseThrow(() -> new NoteNotFoundException(id));

        note.setTitle(updatedNote.getTitle());
        note.setBody(updatedNote.getBody());
        notesRepository.save(note);

        return ResponseEntity.ok().body(note);
    }

    @DeleteMapping(path="/{id}")
    public ResponseEntity<String> deleteNote(@PathVariable long id) {
        notesRepository.deleteById(id);
        return ResponseEntity.status(HttpStatus.NO_CONTENT).build();
    }

}
