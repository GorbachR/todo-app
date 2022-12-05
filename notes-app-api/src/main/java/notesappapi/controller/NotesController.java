package notesappapi.controller;

import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import notesappapi.entity.Notes;
import notesappapi.repository.NotesRepository;


@RestController
@RequestMapping(path = "/api")
public class NotesController {

    private final NotesRepository notesRepository;

    public NotesController(NotesRepository notesRepository) {
        this.notesRepository = notesRepository;
    }
     
    @GetMapping(path = "/notes")
    public List<Notes> getNotes() {
        return notesRepository.findAll(); 
    }
}
