package notesappapi.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import notesappapi.entity.User;

public interface UserRepository extends JpaRepository<User, Long> {
    
}
