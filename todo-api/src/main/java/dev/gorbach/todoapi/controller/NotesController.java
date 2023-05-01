package dev.gorbach.todoapi.controller;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import org.modelmapper.ModelMapper;
import org.modelmapper.PropertyMap;
import org.springframework.boot.autoconfigure.data.web.SpringDataWebProperties.Sort;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;

import dev.gorbach.todoapi.dto.NoteDto;
import dev.gorbach.todoapi.entity.Note;
import dev.gorbach.todoapi.repository.NotesRepository;

@RestController
public class NotesController {

    NotesRepository notesRepository;
    ModelMapper modelMapper;
    
    public NotesController(NotesRepository notesRepository, ModelMapper modelMapper) {

        this.notesRepository = notesRepository;
        this.modelMapper = modelMapper;
    }

    @GetMapping(path = "/notes")
    public ResponseEntity<List<Note>> getAllNotes(Pageable pageable) {

        Page<Note> notes = notesRepository.findAll(pageable);

        return ResponseEntity.ok().body(notes.getContent());
    }

    @PostMapping(path = "/notes")
    public ResponseEntity<Note> saveNote(@RequestBody NoteDto noteDto) {
        Note note = new Note();
        note.setContent(noteDto.getContent());
        note.setCompleted(noteDto.getIsCompleted());

        Note savedNote = notesRepository.save(note);

        return ResponseEntity.status(HttpStatus.CREATED).body(savedNote);
    }

    @PutMapping(path = "/notes/{id}")
    public ResponseEntity<Note> updateNote(@PathVariable UUID id, @RequestBody NoteDto noteDto) {
        Note noteToEdit = notesRepository.findById(id).orElseThrow(() -> 
        new ResponseStatusException(HttpStatus.NOT_FOUND, "Note doesn't exist!"));

        modelMapper.map(noteDto, noteToEdit);

        notesRepository.save(noteToEdit);

        return ResponseEntity.status(HttpStatus.OK).body(noteToEdit);
    }

    @DeleteMapping(path = "notes/{id}")
    public ResponseEntity<Note> deleteNote(@PathVariable UUID id) {
        notesRepository.deleteById(id);
        return ResponseEntity.status(HttpStatus.NO_CONTENT).build();
    }
}   
