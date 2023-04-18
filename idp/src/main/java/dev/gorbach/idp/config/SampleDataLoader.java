package dev.gorbach.idp.config;

import java.util.Arrays;
import java.util.HashSet;

import org.springframework.boot.CommandLineRunner;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Component;

import dev.gorbach.idp.entity.Role;
import dev.gorbach.idp.entity.User;
import dev.gorbach.idp.repository.RoleRepository;
import dev.gorbach.idp.repository.UserRepository;

@Component
public class SampleDataLoader implements CommandLineRunner {

    private final UserRepository userRepository;
    private final RoleRepository roleRepository;
    private final PasswordEncoder passwordEncoder;

    public SampleDataLoader(UserRepository userRepository, RoleRepository roleRepository,
            PasswordEncoder passwordEncoder) {
        this.userRepository = userRepository;
        this.roleRepository = roleRepository;
        this.passwordEncoder = passwordEncoder;
    }

    @Override
    public void run(String... args) throws Exception {
        Role userRole = new Role("USER");
        Role adminRole = new Role("ADMIn");

        roleRepository.saveAll(Arrays.asList(userRole, adminRole));

        User user = new User();
        user.setEmail("raphael_gorbach@gmx.at");
        user.setPassword(passwordEncoder.encode("test"));
        user.setRoles(new HashSet<Role>(roleRepository.findAll()));

        userRepository.save(user);
    }

}