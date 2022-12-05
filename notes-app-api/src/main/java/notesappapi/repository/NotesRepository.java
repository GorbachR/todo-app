package notesappapi.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import notesappapi.entity.Notes;

public interface NotesRepository extends JpaRepository<Notes, Long> {
    
}
