package notesappapi.config;

import java.util.List;
import java.util.concurrent.ThreadLocalRandom;
import java.util.stream.IntStream;
import org.springframework.boot.CommandLineRunner;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Component;
import com.github.javafaker.Faker;
import notesappapi.entity.Note;
import notesappapi.entity.User;
import notesappapi.repository.NotesRepository;
import notesappapi.repository.UserRepository;

@Component
public class SampleDataLoader implements CommandLineRunner {

    private final NotesRepository notesRepository;
    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final Faker faker;

    public SampleDataLoader(NotesRepository notesRepository, UserRepository userRepository,
            PasswordEncoder passwordEncoder) {
        this.notesRepository = notesRepository;
        this.userRepository = userRepository;
        this.passwordEncoder = passwordEncoder;
        this.faker = new Faker();
    }

    @Override
    public void run(String... args) throws Exception {

        User user = new User();
        user.setEmail("gorbach");
        user.setPassword(passwordEncoder.encode("password"));
        userRepository.save(user);

        List<Note> fakeNotes = IntStream.range(0, 100).mapToObj((i) -> {
            Note note = new Note();

            note.setUser(userRepository.findById(ThreadLocalRandom.current().nextLong(1, userRepository.count() + 1))
                    .orElseThrow(() -> new RuntimeException()));

            note.setTitle(faker.animal().name());
            note.setBody(faker.lorem().paragraph());
            return note;
        }).toList();

        notesRepository.saveAll(fakeNotes);
    }

}
