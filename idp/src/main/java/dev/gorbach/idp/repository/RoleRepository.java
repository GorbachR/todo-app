package dev.gorbach.idp.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import dev.gorbach.idp.entity.Role;

public interface RoleRepository extends JpaRepository<Role, Long> {

}