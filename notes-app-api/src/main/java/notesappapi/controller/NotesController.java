package notesappapi.controller;

import java.util.List;

import org.modelmapper.ModelMapper;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
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
import jakarta.validation.Valid;
import notesappapi.entity.Note;
import notesappapi.exception.NoteNotFoundException;
import notesappapi.model.NoteDto;
import notesappapi.repository.NotesRepository;

@RestController
@RequestMapping(path = "/api/notes")
public class NotesController {

    private final NotesRepository notesRepository;

    private final ModelMapper modelMapper;

    public NotesController(NotesRepository notesRepository, ModelMapper modelMapper) {
        this.notesRepository = notesRepository;
        this.modelMapper = modelMapper;
    }

    @GetMapping
    public ResponseEntity<List<NoteDto>> getNotes(@RequestParam int page, @RequestParam(defaultValue = "3") int size) {

        Pageable currentPage = PageRequest.of(page, size);
        Page<NoteDto> result = notesRepository.findBy(currentPage, NoteDto.class);

        return ResponseEntity.ok().body(result.getContent());
    }

    @PostMapping(consumes = { "application/json" })
    public ResponseEntity<NoteDto> addNote(@Valid @RequestBody NoteDto newNote) {

        Note note = convertToEntity(newNote);
        notesRepository.save(note);
        return ResponseEntity.status(HttpStatus.CREATED).body(convertToDto(note));
    }

    @PutMapping(path = "/{id}", consumes = { "application/json" })
    public ResponseEntity<NoteDto> editNote(@PathVariable long id, @Valid @RequestBody NoteDto updatedNote) {

        Note note = notesRepository.findById(id)
                .orElseThrow(() -> new NoteNotFoundException(id));

        note.setTitle(updatedNote.getTitle());
        note.setBody(updatedNote.getBody());
        notesRepository.save(note);

        return ResponseEntity.ok().body(convertToDto(note));
    }

    @DeleteMapping(path = "/{id}")
    public ResponseEntity<String> deleteNote(@PathVariable long id) {
        notesRepository.deleteById(id);
        return ResponseEntity.status(HttpStatus.NO_CONTENT).build();
    }

    private NoteDto convertToDto(Note note) {
        return modelMapper.map(note, NoteDto.class);
    }

    private Note convertToEntity(NoteDto noteDto) {
        return modelMapper.map(noteDto, Note.class);
    }

}