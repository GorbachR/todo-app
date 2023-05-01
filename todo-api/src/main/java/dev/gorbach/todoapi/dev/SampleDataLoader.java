package dev.gorbach.todoapi.dev;

import java.util.List;
import java.util.stream.IntStream;

import org.springframework.boot.CommandLineRunner;
import org.springframework.context.annotation.Profile;
import org.springframework.stereotype.Component;

import com.github.javafaker.Faker;

import dev.gorbach.todoapi.entity.Note;
import dev.gorbach.todoapi.repository.NotesRepository;

@Component
@Profile("dev")
public class SampleDataLoader implements CommandLineRunner {

    private final NotesRepository notesRepository;
    private final Faker faker;

    public SampleDataLoader(NotesRepository notesRepository) {
        this.notesRepository = notesRepository;
        this.faker = new Faker();
    }

    @Override
    public void run(String... args) throws Exception {
        List<Note> notes = IntStream.rangeClosed(0, 100)
            .mapToObj((i) -> {
                Note note = new Note();
                note.setContent(faker.cat().name());
                note.setCompleted(faker.bool().bool());
                return note;
            }).toList();

        notesRepository.saveAll(notes);
    }
    
}