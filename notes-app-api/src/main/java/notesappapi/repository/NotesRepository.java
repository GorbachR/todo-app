package notesappapi.repository;


import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.List;
import notesappapi.entity.Note;

public interface NotesRepository extends JpaRepository<Note, Long> {
   public List<Note> findByIdLessThanEqualOrderByIdDesc(long id, Pageable pageable);
}
