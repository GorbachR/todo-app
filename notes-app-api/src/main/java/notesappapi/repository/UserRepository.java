package notesappapi.repository;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import notesappapi.entity.User;

public interface UserRepository extends JpaRepository<User, Long> {
   public Optional<User> findByEmail(String email);
}