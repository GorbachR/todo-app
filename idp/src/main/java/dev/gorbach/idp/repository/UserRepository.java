package dev.gorbach.idp.repository;

import java.util.Optional;
import org.springframework.data.jpa.repository.JpaRepository;
import dev.gorbach.idp.entity.User;

public interface UserRepository extends JpaRepository<User, Long> {
    Optional<User> findByEmail(String email);
}