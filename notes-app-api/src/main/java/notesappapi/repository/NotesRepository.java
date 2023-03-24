package notesappapi.repository;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import notesappapi.entity.Note;

public interface NotesRepository extends JpaRepository<Note, Long> {
   <T> Page<T> findBy(Pageable pageable, Class<T> type);
}