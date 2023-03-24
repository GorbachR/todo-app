package notesappapi.controller;

import java.util.Optional;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import notesappapi.entity.User;
import notesappapi.model.UserData;
import notesappapi.repository.UserRepository;
import notesappapi.service.JwtService;

@RestController
@RequestMapping("/auth")
public class AuthController {

    private UserRepository userRepository;

    private PasswordEncoder passwordEncoder;

    private JwtService jwtService;

    private AuthenticationManager authManager;

    public AuthController(UserRepository userRepository,
            PasswordEncoder passwordEncoder,
            JwtService jwtService,
            AuthenticationManager authManager) {
        this.userRepository = userRepository;
        this.passwordEncoder = passwordEncoder;
        this.jwtService = jwtService;
        this.authManager = authManager;
    }

    @PostMapping(path = "/register", consumes = { "application/json" })
    ResponseEntity<User> register(@RequestBody UserData inputUser) {

        Optional<User> optionalUser = userRepository.findByEmail(inputUser.username());

        if (optionalUser.isPresent()) {
            return ResponseEntity.badRequest().build();
        }

        User user = new User();
        user.setEmail(inputUser.username());
        user.setPassword(passwordEncoder.encode(inputUser.password()));

        User created = userRepository.save(user);

        return ResponseEntity.ok(created);
    }

    @PostMapping(path = "/login", consumes = { "application/json" })
    ResponseEntity<String> login(@RequestBody UserData loginData) {

        Authentication authentication = authManager
                .authenticate(new UsernamePasswordAuthenticationToken(loginData.username(), loginData.password()));
        return ResponseEntity.ok().body(jwtService.generateToken(authentication));
    }
}
